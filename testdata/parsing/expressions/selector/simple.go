package main
var e = a.b

/**-----
Go file
  PackageDeclaration(main)
    PsiElement(KEYWORD_PACKAGE)('package')
    PsiWhiteSpace(' ')
    PsiElement(IDENTIFIER)('main')
  PsiWhiteSpace('\n')
  VarDeclarationsImpl
    PsiElement(KEYWORD_VAR)('var')
    PsiWhiteSpace(' ')
    VarDeclarationImpl
      LiteralIdentifierImpl
        PsiElement(IDENTIFIER)('e')
      PsiWhiteSpace(' ')
      PsiElement(=)('=')
      PsiWhiteSpace(' ')
      SelectorExpressionImpl
        LiteralExpressionImpl
          LiteralIdentifierImpl
            PsiElement(IDENTIFIER)('a')
        PsiElement(.)('.')
        LiteralIdentifierImpl
          PsiElement(IDENTIFIER)('b')
  PsiWhiteSpace('\n')
