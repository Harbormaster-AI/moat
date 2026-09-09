import React, { Component } from 'react'
import ConditionService from '../services/ConditionService'

class ListConditionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                conditions: []
        }
        this.addCondition = this.addCondition.bind(this);
        this.editCondition = this.editCondition.bind(this);
        this.deleteCondition = this.deleteCondition.bind(this);
    }

    deleteCondition(id){
        ConditionService.deleteCondition(id).then( res => {
            this.setState({conditions: this.state.conditions.filter(condition => condition.conditionId !== id)});
        });
    }
    viewCondition(id){
        this.props.history.push(`/view-condition/${id}`);
    }
    editCondition(id){
        this.props.history.push(`/add-condition/${id}`);
    }

    componentDidMount(){
        ConditionService.getConditions().then((res) => {
            this.setState({ conditions: res.data});
        });
    }

    addCondition(){
        this.props.history.push('/add-condition/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Condition List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCondition}> Add Condition</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Code </th>
                                    <th> OnsetDate </th>
                                    <th> AbatementDate </th>
                                    <th> ClinicalStatus </th>
                                    <th> VerificationStatus </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.conditions.map(
                                        condition => 
                                        <tr key = {condition.conditionId}>
                                             <td> { condition.code } </td>
                                             <td> { condition.onsetDate } </td>
                                             <td> { condition.abatementDate } </td>
                                             <td> { condition.clinicalStatus } </td>
                                             <td> { condition.verificationStatus } </td>
                                             <td>
                                                 <button onClick={ () => this.editCondition(condition.conditionId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCondition(condition.conditionId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCondition(condition.conditionId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListConditionComponent
