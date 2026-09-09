import React, { Component } from 'react'
import CorrectiveActionService from '../services/CorrectiveActionService'

class ListCorrectiveActionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                correctiveActions: []
        }
        this.addCorrectiveAction = this.addCorrectiveAction.bind(this);
        this.editCorrectiveAction = this.editCorrectiveAction.bind(this);
        this.deleteCorrectiveAction = this.deleteCorrectiveAction.bind(this);
    }

    deleteCorrectiveAction(id){
        CorrectiveActionService.deleteCorrectiveAction(id).then( res => {
            this.setState({correctiveActions: this.state.correctiveActions.filter(correctiveAction => correctiveAction.correctiveActionId !== id)});
        });
    }
    viewCorrectiveAction(id){
        this.props.history.push(`/view-correctiveAction/${id}`);
    }
    editCorrectiveAction(id){
        this.props.history.push(`/add-correctiveAction/${id}`);
    }

    componentDidMount(){
        CorrectiveActionService.getCorrectiveActions().then((res) => {
            this.setState({ correctiveActions: res.data});
        });
    }

    addCorrectiveAction(){
        this.props.history.push('/add-correctiveAction/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">CorrectiveAction List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCorrectiveAction}> Add CorrectiveAction</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> CapaNumber </th>
                                    <th> RootCause </th>
                                    <th> CorrectiveAction </th>
                                    <th> VerificationDate </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.correctiveActions.map(
                                        correctiveAction => 
                                        <tr key = {correctiveAction.correctiveActionId}>
                                             <td> { correctiveAction.capaNumber } </td>
                                             <td> { correctiveAction.rootCause } </td>
                                             <td> { correctiveAction.correctiveAction } </td>
                                             <td> { correctiveAction.verificationDate } </td>
                                             <td> { correctiveAction.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editCorrectiveAction(correctiveAction.correctiveActionId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCorrectiveAction(correctiveAction.correctiveActionId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCorrectiveAction(correctiveAction.correctiveActionId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListCorrectiveActionComponent
