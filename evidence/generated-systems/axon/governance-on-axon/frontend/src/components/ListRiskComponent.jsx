import React, { Component } from 'react'
import RiskService from '../services/RiskService'

class ListRiskComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                risks: []
        }
        this.addRisk = this.addRisk.bind(this);
        this.editRisk = this.editRisk.bind(this);
        this.deleteRisk = this.deleteRisk.bind(this);
    }

    deleteRisk(id){
        RiskService.deleteRisk(id).then( res => {
            this.setState({risks: this.state.risks.filter(risk => risk.riskId !== id)});
        });
    }
    viewRisk(id){
        this.props.history.push(`/view-risk/${id}`);
    }
    editRisk(id){
        this.props.history.push(`/add-risk/${id}`);
    }

    componentDidMount(){
        RiskService.getRisks().then((res) => {
            this.setState({ risks: res.data});
        });
    }

    addRisk(){
        this.props.history.push('/add-risk/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Risk List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addRisk}> Add Risk</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Description </th>
                                    <th> InherentRiskScore </th>
                                    <th> ResidualRiskScore </th>
                                    <th> Category </th>
                                    <th> Impact </th>
                                    <th> Likelihood </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.risks.map(
                                        risk => 
                                        <tr key = {risk.riskId}>
                                             <td> { risk.name } </td>
                                             <td> { risk.description } </td>
                                             <td> { risk.inherentRiskScore } </td>
                                             <td> { risk.residualRiskScore } </td>
                                             <td> { risk.category } </td>
                                             <td> { risk.impact } </td>
                                             <td> { risk.likelihood } </td>
                                             <td> { risk.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editRisk(risk.riskId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteRisk(risk.riskId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewRisk(risk.riskId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListRiskComponent
