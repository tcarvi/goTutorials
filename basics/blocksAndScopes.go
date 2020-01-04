package basics

/*
 * Há blocos explícitos e blocos implícitos (cada bloco apresenta escopo para as suas variáveis)
 * Blocos implícitos:
 *   Há o "universe block" que compreende todo o código fonte que é indicado ou importado na package main.
 *   Há o "package block" que compreende todo o código fonte de uma package.
 *   Há o "file block" que compreende todo o código fonte de um arquivo.
 *   Há "implicit block" para cada "if", "for", e "switch" statements.
 *   Há "implicit block" para cada clause no "switch" ou no "select".
 * Blocos explícitos:
 *   Uma sequência de "declarations" e "statements", formalmente indicada e limitada por { }
 *   Sua sequência de "declarations" e "statements" pode estar vazia
 *   Pode haver hierarquias internas de blocos
 *   Block = "{" StatementList "}"
 *	     StatementList = { Statement ";" }
 */

// TBlocksAndScopes resume https://golang.org/ref/spec#Blocks e https://golang.org/ref/spec#Declarations_and_scope
func TBlocksAndScopes() {
	// Não se pode duplicar a declaração de identifiers, ao mesmo tempo, em "file block" e em "package block".
	// The scope of a declared identifier is the extent of source text in which the identifier denotes the specified constant, type, variable, function, label, or package.

	// Escopos:
	// predeclared identifier ->  universe block
	// identifier denoting a constant, type, variable, or function (but not method) declared at top level (outside any function) -> package block
	// package name of an imported package -> file block of the file containing the import declaration
	// identifier denoting a method receiver, function parameter, or result variable -> function body
	// constant or variable identifier declared inside a function -> begins at the end of the ConstSpec or VarSpec (ShortVarDecl for short variable declarations) and ends at the end of the innermost containing block
	// type identifier declared inside a function -> begins at the identifier in the TypeSpec and ends at the end of the innermost containing block

	// Redeclaração de variáveis
	// - identifier declared in a block may be redeclared in an inner block.
	// - While the identifier of the inner declaration is in scope, it denotes the entity declared by the inner declaration.

	// The package clause is not a declaration; the package name does not appear in any scope. Its purpose is to identify the files belonging to the same package and to specify the default package name for import declarations.

}
