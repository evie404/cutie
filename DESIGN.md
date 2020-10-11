# Design of cutie

## Separation of Concern

The core motivation behind cutie is that models at every layer should be distinct and each suited for its requirements and use-cases. We consider models in three different contexts:

- Database persistence
- Application logic
- External interfaces

While it is tempting to use the same model across all context, we might run into different issues. For example, suppose we define a protobuf message `User` and use it across all layers:

- There may be fields we'd like to persist without exposing them to external interfaces. For example, we would store passwords but never expose the stored content, even encrypted, to even the users themselves.
- Conversely, there may be fields we expose to external interfaces without storing them in the database. For example, we might want to return a user's initials by composing first and last names without actually storing the initials.
- Each layer has varying preferences of how to expose the same property. For example, when it comes to nullness, database types prefer `sql.NullInt64` types whereas protobuf prefers `*int64` for proto2 and `google.protobuf.Int64Value` for proto3.
- Often times, application logic requires juggling variables of similar types. An `int64` type may be an ID reference to a user or an ID reference to an external user (think third-party links like Twitter, Facebook, etc.). It can be beneficial to create named types like `UserID` to differentiate between IDs of different domains to prevent mistakes in development of passing the wrong ids across pieces of logic.

Addressing all of the above concerns, we might use three different struct types to represent each layer in our application:

- A database model `UserRow` that makes use of `sql.NullInt64` types
- An application model `user.User` that makes use of named types like `*UserID` instead of `*int64`
- Protobuf message `proto_user.User` for exposing the model to external users

cutie currently supports the generation of the database models from PostgresSQL schemas with [sqlc](https://github.com/kyleconroy/sqlc) along with mocks with [gomock](https://github.com/golang/mock). At present, cutie expects users to supply their own application and API models. The future plan of cutie is to also generate protobuf messages for API and integrate with the models from the other layers. The opportunities and roadblocks are described in the next section.

## Model Transformations

Separation of Concern seems like a decent idea until we realize for each domain object, not only do we have to create three different representations, we also have to create transformations between them. This compounds for each new field and model added to our application.

Fortunately, the transformations between individual fields is quite standardized and repetitive. Transforming between `sql.NullInt64` to `*UserID` to `google.protobuf.Int64Value` is the same regardless of whether this user ID is embedded as part of the user domain object or say purchases, which contains such user ID reference for the buyer. Furthermore, we can infer that the translations to `*PurchaseID` is no different from that of `*UserID` assuming they share the same underlying value type of `int64`.

This implies that we can produce standardized transformations between individual fields from inferring their types. Should this work, this might mean that once we define above models, we can automatically generate the transformations between them, avoiding the need of repetition and the potential of human error. We might even be able to generate unit tests for these transforms.

That is the aspiration at least. Having these transforms work without costly runtime reflections means we have to understand the types before compilation to generate the appropriate code for the compiler. At the same time, the only way to really understand the types is through the compiler. This chicken-and-egg situation might be alleviated with dynamic loading, which Go doesn't allow for good reasons. This means unless we want users to import cutie to their models (or the other way around), cutie needs to use AST to infer the types. The inference is more lossy than direct reflections through the compiler, but it potentially allows cutie to be used as a standalone CLI.
