import React, { Component } from 'react'
import KPIService from '../services/KPIService'

class ListKPIComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                kPIs: []
        }
        this.addKPI = this.addKPI.bind(this);
        this.editKPI = this.editKPI.bind(this);
        this.deleteKPI = this.deleteKPI.bind(this);
    }

    deleteKPI(id){
        KPIService.deleteKPI(id).then( res => {
            this.setState({kPIs: this.state.kPIs.filter(kPI => kPI.kPIId !== id)});
        });
    }
    viewKPI(id){
        this.props.history.push(`/view-kPI/${id}`);
    }
    editKPI(id){
        this.props.history.push(`/add-kPI/${id}`);
    }

    componentDidMount(){
        KPIService.getKPIs().then((res) => {
            this.setState({ kPIs: res.data});
        });
    }

    addKPI(){
        this.props.history.push('/add-kPI/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">KPI List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addKPI}> Add KPI</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> TargetValue </th>
                                    <th> MetricType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.kPIs.map(
                                        kPI => 
                                        <tr key = {kPI.kPIId}>
                                             <td> { kPI.targetValue } </td>
                                             <td> { kPI.metricType } </td>
                                             <td>
                                                 <button onClick={ () => this.editKPI(kPI.kPIId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteKPI(kPI.kPIId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewKPI(kPI.kPIId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListKPIComponent
