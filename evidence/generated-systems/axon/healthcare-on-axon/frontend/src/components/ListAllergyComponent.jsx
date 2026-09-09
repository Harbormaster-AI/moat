import React, { Component } from 'react'
import AllergyService from '../services/AllergyService'

class ListAllergyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                allergys: []
        }
        this.addAllergy = this.addAllergy.bind(this);
        this.editAllergy = this.editAllergy.bind(this);
        this.deleteAllergy = this.deleteAllergy.bind(this);
    }

    deleteAllergy(id){
        AllergyService.deleteAllergy(id).then( res => {
            this.setState({allergys: this.state.allergys.filter(allergy => allergy.allergyId !== id)});
        });
    }
    viewAllergy(id){
        this.props.history.push(`/view-allergy/${id}`);
    }
    editAllergy(id){
        this.props.history.push(`/add-allergy/${id}`);
    }

    componentDidMount(){
        AllergyService.getAllergys().then((res) => {
            this.setState({ allergys: res.data});
        });
    }

    addAllergy(){
        this.props.history.push('/add-allergy/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Allergy List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAllergy}> Add Allergy</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Substance </th>
                                    <th> Reaction </th>
                                    <th> Severity </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.allergys.map(
                                        allergy => 
                                        <tr key = {allergy.allergyId}>
                                             <td> { allergy.substance } </td>
                                             <td> { allergy.reaction } </td>
                                             <td> { allergy.severity } </td>
                                             <td> { allergy.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editAllergy(allergy.allergyId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAllergy(allergy.allergyId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAllergy(allergy.allergyId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAllergyComponent
