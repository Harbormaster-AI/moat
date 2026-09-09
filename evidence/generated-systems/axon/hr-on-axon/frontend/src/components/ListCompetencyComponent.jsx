import React, { Component } from 'react'
import CompetencyService from '../services/CompetencyService'

class ListCompetencyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                competencys: []
        }
        this.addCompetency = this.addCompetency.bind(this);
        this.editCompetency = this.editCompetency.bind(this);
        this.deleteCompetency = this.deleteCompetency.bind(this);
    }

    deleteCompetency(id){
        CompetencyService.deleteCompetency(id).then( res => {
            this.setState({competencys: this.state.competencys.filter(competency => competency.competencyId !== id)});
        });
    }
    viewCompetency(id){
        this.props.history.push(`/view-competency/${id}`);
    }
    editCompetency(id){
        this.props.history.push(`/add-competency/${id}`);
    }

    componentDidMount(){
        CompetencyService.getCompetencys().then((res) => {
            this.setState({ competencys: res.data});
        });
    }

    addCompetency(){
        this.props.history.push('/add-competency/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Competency List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCompetency}> Add Competency</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Category </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.competencys.map(
                                        competency => 
                                        <tr key = {competency.competencyId}>
                                             <td> { competency.name } </td>
                                             <td> { competency.category } </td>
                                             <td>
                                                 <button onClick={ () => this.editCompetency(competency.competencyId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCompetency(competency.competencyId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCompetency(competency.competencyId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListCompetencyComponent
