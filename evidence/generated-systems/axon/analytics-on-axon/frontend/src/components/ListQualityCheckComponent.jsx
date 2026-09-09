import React, { Component } from 'react'
import QualityCheckService from '../services/QualityCheckService'

class ListQualityCheckComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                qualityChecks: []
        }
        this.addQualityCheck = this.addQualityCheck.bind(this);
        this.editQualityCheck = this.editQualityCheck.bind(this);
        this.deleteQualityCheck = this.deleteQualityCheck.bind(this);
    }

    deleteQualityCheck(id){
        QualityCheckService.deleteQualityCheck(id).then( res => {
            this.setState({qualityChecks: this.state.qualityChecks.filter(qualityCheck => qualityCheck.qualityCheckId !== id)});
        });
    }
    viewQualityCheck(id){
        this.props.history.push(`/view-qualityCheck/${id}`);
    }
    editQualityCheck(id){
        this.props.history.push(`/add-qualityCheck/${id}`);
    }

    componentDidMount(){
        QualityCheckService.getQualityChecks().then((res) => {
            this.setState({ qualityChecks: res.data});
        });
    }

    addQualityCheck(){
        this.props.history.push('/add-qualityCheck/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">QualityCheck List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addQualityCheck}> Add QualityCheck</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> CheckedAt </th>
                                    <th> ObservedValue </th>
                                    <th> SampleSize </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.qualityChecks.map(
                                        qualityCheck => 
                                        <tr key = {qualityCheck.qualityCheckId}>
                                             <td> { qualityCheck.checkedAt } </td>
                                             <td> { qualityCheck.observedValue } </td>
                                             <td> { qualityCheck.sampleSize } </td>
                                             <td> { qualityCheck.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editQualityCheck(qualityCheck.qualityCheckId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteQualityCheck(qualityCheck.qualityCheckId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewQualityCheck(qualityCheck.qualityCheckId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListQualityCheckComponent
