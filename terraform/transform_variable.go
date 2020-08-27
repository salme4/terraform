package terraform

import (
	"fmt"

	"github.com/hashicorp/terraform/addrs"
	"github.com/hashicorp/terraform/configs"
)

// RootVariableTransformer is a GraphTransformer that adds all the root
// variables to the graph.
//
// Root variables are currently no-ops but they must be added to the
// graph since downstream things that depend on them must be able to
// reach them.
type RootVariableTransformer struct {
	Config *configs.Config
}

func (t *RootVariableTransformer) Transform(g *Graph) error {
	// We can have no variables if we have no config.
	if t.Config == nil {
		return nil
	}

	// We're only considering root module variables here, since child
	// module variables are handled by ModuleVariableTransformer.
	vars := t.Config.Module.Variables

	fmt.Println("Adding variables")
	// Add all variables here
	for _, v := range vars {
		fmt.Printf("Adding NodeRootVariable: sensitivity for %s is %v\n", v.Name, v.Sensitive)
		node := &NodeRootVariable{
			Addr: addrs.InputVariable{
				Name:      v.Name,
				Sensitive: true, //v.Sensitive,
			},
			Config: v,
		}
		g.Add(node)
	}

	return nil
}
