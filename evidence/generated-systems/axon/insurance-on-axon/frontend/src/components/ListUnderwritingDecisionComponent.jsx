import React, { Component } from 'react'
import UnderwritingDecisionService from '../services/UnderwritingDecisionService'

class ListUnderwritingDecisionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                underwritingDecisions: []
        }
        this.addUnderwritingDecision = this.addUnderwritingDecision.bind(this);
        this.editUnderwritingDecision = this.editUnderwritingDecision.bind(this);
        this.deleteUnderwritingDecision = this.deleteUnderwritingDecision.bind(this);
    }

    deleteUnderwritingDecision(id){
        UnderwritingDecisionService.deleteUnderwritingDecision(id).then( res => {
            this.setState({underwritingDecisions: this.state.underwritingDecisions.filter(underwritingDecision => underwritingDecision.underwritingDecisionId !== id)});
        });
    }
    viewUnderwritingDecision(id){
        this.props.history.push(`/view-underwritingDecision/${id}`);
    }
    editUnderwritingDecision(id){
        this.props.history.push(`/add-underwritingDecision/${id}`);
    }

    componentDidMount(){
        UnderwritingDecisionService.getUnderwritingDecisions().then((res) => {
            this.setState({ underwritingDecisions: res.data});
        });
    }

    addUnderwritingDecision(){
        this.props.history.push('/add-underwritingDecision/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">UnderwritingDecision List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addUnderwritingDecision}> Add UnderwritingDecision</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Notes </th>
                                    <th> DecisionDate </th>
                                    <th> Decision </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.underwritingDecisions.map(
                                        underwritingDecision => 
                                        <tr key = {underwritingDecision.underwritingDecisionId}>
                                             <td> { underwritingDecision.notes } </td>
                                             <td> { underwritingDecision.decisionDate } </td>
                                             <td> { underwritingDecision.decision } </td>
                                             <td>
                                                 <button onClick={ () => this.editUnderwritingDecision(underwritingDecision.underwritingDecisionId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteUnderwritingDecision(underwritingDecision.underwritingDecisionId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewUnderwritingDecision(underwritingDecision.underwritingDecisionId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListUnderwritingDecisionComponent
