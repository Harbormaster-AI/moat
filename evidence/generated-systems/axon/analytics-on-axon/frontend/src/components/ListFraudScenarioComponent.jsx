import React, { Component } from 'react'
import FraudScenarioService from '../services/FraudScenarioService'

class ListFraudScenarioComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                fraudScenarios: []
        }
        this.addFraudScenario = this.addFraudScenario.bind(this);
        this.editFraudScenario = this.editFraudScenario.bind(this);
        this.deleteFraudScenario = this.deleteFraudScenario.bind(this);
    }

    deleteFraudScenario(id){
        FraudScenarioService.deleteFraudScenario(id).then( res => {
            this.setState({fraudScenarios: this.state.fraudScenarios.filter(fraudScenario => fraudScenario.fraudScenarioId !== id)});
        });
    }
    viewFraudScenario(id){
        this.props.history.push(`/view-fraudScenario/${id}`);
    }
    editFraudScenario(id){
        this.props.history.push(`/add-fraudScenario/${id}`);
    }

    componentDidMount(){
        FraudScenarioService.getFraudScenarios().then((res) => {
            this.setState({ fraudScenarios: res.data});
        });
    }

    addFraudScenario(){
        this.props.history.push('/add-fraudScenario/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">FraudScenario List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addFraudScenario}> Add FraudScenario</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> RiskAppetite </th>
                                    <th> DetectionType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.fraudScenarios.map(
                                        fraudScenario => 
                                        <tr key = {fraudScenario.fraudScenarioId}>
                                             <td> { fraudScenario.name } </td>
                                             <td> { fraudScenario.riskAppetite } </td>
                                             <td> { fraudScenario.detectionType } </td>
                                             <td>
                                                 <button onClick={ () => this.editFraudScenario(fraudScenario.fraudScenarioId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteFraudScenario(fraudScenario.fraudScenarioId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewFraudScenario(fraudScenario.fraudScenarioId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListFraudScenarioComponent
