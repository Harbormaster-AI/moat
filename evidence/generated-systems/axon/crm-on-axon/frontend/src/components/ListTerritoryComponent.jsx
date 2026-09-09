import React, { Component } from 'react'
import TerritoryService from '../services/TerritoryService'

class ListTerritoryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                territorys: []
        }
        this.addTerritory = this.addTerritory.bind(this);
        this.editTerritory = this.editTerritory.bind(this);
        this.deleteTerritory = this.deleteTerritory.bind(this);
    }

    deleteTerritory(id){
        TerritoryService.deleteTerritory(id).then( res => {
            this.setState({territorys: this.state.territorys.filter(territory => territory.territoryId !== id)});
        });
    }
    viewTerritory(id){
        this.props.history.push(`/view-territory/${id}`);
    }
    editTerritory(id){
        this.props.history.push(`/add-territory/${id}`);
    }

    componentDidMount(){
        TerritoryService.getTerritorys().then((res) => {
            this.setState({ territorys: res.data});
        });
    }

    addTerritory(){
        this.props.history.push('/add-territory/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Territory List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addTerritory}> Add Territory</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Region </th>
                                    <th> TerritoryType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.territorys.map(
                                        territory => 
                                        <tr key = {territory.territoryId}>
                                             <td> { territory.name } </td>
                                             <td> { territory.region } </td>
                                             <td> { territory.territoryType } </td>
                                             <td>
                                                 <button onClick={ () => this.editTerritory(territory.territoryId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteTerritory(territory.territoryId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewTerritory(territory.territoryId)} className="btn btn-outline-info btn-sm">View </button>
                                             </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>

                 </div>

            </div>
        )
    }
}

export default ListTerritoryComponent
