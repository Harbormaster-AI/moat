import React, { Component } from 'react'
import HealthSystemService from '../services/HealthSystemService'

class ListHealthSystemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                healthSystems: []
        }
        this.addHealthSystem = this.addHealthSystem.bind(this);
        this.editHealthSystem = this.editHealthSystem.bind(this);
        this.deleteHealthSystem = this.deleteHealthSystem.bind(this);
    }

    deleteHealthSystem(id){
        HealthSystemService.deleteHealthSystem(id).then( res => {
            this.setState({healthSystems: this.state.healthSystems.filter(healthSystem => healthSystem.healthSystemId !== id)});
        });
    }
    viewHealthSystem(id){
        this.props.history.push(`/view-healthSystem/${id}`);
    }
    editHealthSystem(id){
        this.props.history.push(`/add-healthSystem/${id}`);
    }

    componentDidMount(){
        HealthSystemService.getHealthSystems().then((res) => {
            this.setState({ healthSystems: res.data});
        });
    }

    addHealthSystem(){
        this.props.history.push('/add-healthSystem/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">HealthSystem List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addHealthSystem}> Add HealthSystem</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> LegalName </th>
                                    <th> HeadquartersCountry </th>
                                    <th> Website </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.healthSystems.map(
                                        healthSystem => 
                                        <tr key = {healthSystem.healthSystemId}>
                                             <td> { healthSystem.name } </td>
                                             <td> { healthSystem.legalName } </td>
                                             <td> { healthSystem.headquartersCountry } </td>
                                             <td> { healthSystem.website } </td>
                                             <td>
                                                 <button onClick={ () => this.editHealthSystem(healthSystem.healthSystemId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteHealthSystem(healthSystem.healthSystemId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewHealthSystem(healthSystem.healthSystemId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListHealthSystemComponent
