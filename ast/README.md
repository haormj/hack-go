## ast

### Usage

```shell
go get -u github.com/lu4p/astextract/cmd/astextract
astextract main.go
```

### Result

```
&ast.File {
  Package: 1,
  Name: &ast.Ident {
    Name: "main",
  },
  Decls: []ast.Decl {
    &ast.GenDecl {
      Tok: token.IMPORT,
      Specs: []ast.Spec {
        &ast.ImportSpec {
          Path: &ast.BasicLit {
            Kind: token.STRING,
            Value: "\"fmt\"",
          },
        },
      },
    },
    &ast.FuncDecl {
      Name: &ast.Ident {
        Name: "main",
      },
      Type: &ast.FuncType {
        Params: &ast.FieldList {},
      },
      Body: &ast.BlockStmt {
        List: []ast.Stmt {
          &ast.ExprStmt {
            X: &ast.CallExpr {
              Fun: &ast.SelectorExpr {
                X: &ast.Ident {
                  Name: "fmt",
                },
                Sel: &ast.Ident {
                  Name: "Println",
                },
              },
              Args: []ast.Expr {
                &ast.BasicLit {
                  Kind: token.STRING,
                  Value: "\"hello world\"",
                },
              },
            },
          },
        },
      },
    },
  },
  Imports: []*ast.ImportSpec {
    &ast.ImportSpec {
      Path: &ast.BasicLit {
        Kind: token.STRING,
        Value: "\"fmt\"",
      },
    },
  },
}
```
