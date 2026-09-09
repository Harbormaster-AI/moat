import React, { Component } from 'react'
import AnomalyService from '../services/AnomalyService'

class ListAnomalyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                anomalys: []
        }
        this.addAnomaly = this.addAnomaly.bind(this);
        this.editAnomaly = this.editAnomaly.bind(this);
        this.deleteAnomaly = this.deleteAnomaly.bind(this);
    }

    deleteAnomaly(id){
        AnomalyService.deleteAnomaly(id).then( res => {
            this.setState({anomalys: this.state.anomalys.filter(anomaly => anomaly.anomalyId !== id)});
        });
    }
    viewAnomaly(id){
        this.props.history.push(`/view-anomaly/${id}`);
    }
    editAnomaly(id){
        this.props.history.push(`/add-anomaly/${id}`);
    }

    componentDidMount(){
        AnomalyService.getAnomalys().then((res) => {
            this.setState({ anomalys: res.data});
        });
    }

    addAnomaly(){
        this.props.history.push('/add-anomaly/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Anomaly List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAnomaly}> Add Anomaly</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> OccurredAt </th>
                                    <th> Details </th>
                                    <th> AnomalyType </th>
                                    <th> Severity </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.anomalys.map(
                                        anomaly => 
                                        <tr key = {anomaly.anomalyId}>
                                             <td> { anomaly.occurredAt } </td>
                                             <td> { anomaly.details } </td>
                                             <td> { anomaly.anomalyType } </td>
                                             <td> { anomaly.severity } </td>
                                             <td>
                                                 <button onClick={ () => this.editAnomaly(anomaly.anomalyId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAnomaly(anomaly.anomalyId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAnomaly(anomaly.anomalyId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAnomalyComponent
