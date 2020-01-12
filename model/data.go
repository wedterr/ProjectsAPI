package model

type authorsCellection []Author

var AuthorsCellection = authorsCellection{
	Author{
		"1",
		"John",
		"Doe",
		projectsCollection1,
	},
	Author{
		"2",
		"Peter",
		"Parker",
		projectsCollection2,
	},
	Author{
		"3",
		"Sherlock",
		"Holmes",
		projectsCollection3,
	},
}

var projectsCollection1 = ProjectsCollection{
	Project{
		"1",
		"John",
		"Doe",
		"GridStore",
		"Store of grid data",
	},
	Project{
		"2",
		"John",
		"Doe",
		"SeismicStore",
		"Store of seismic data",
	},
}

var projectsCollection2 = ProjectsCollection{
	Project{
		"3",
		"Peter",
		"Parker",
		"StoneStore",
		"Store of stone data",
	},
	Project{
		"4",
		"Peter",
		"Parker",
		"WellsStore",
		"Store of wells data",
	},
}

var projectsCollection3 = ProjectsCollection{
	Project{
		"5",
		"Sherlock",
		"Holmes",
		"BookStore",
		"Book store",
	},
}
