import React, { Component } from 'react'
import Case_Service from '../services/Case_Service'

class ListCase_Component extends Component {
    constructor(props) {
        super(props)

        this.state = {
                case_s: []
        }
        this.addCase_ = this.addCase_.bind(this);
        this.editCase_ = this.editCase_.bind(this);
        this.deleteCase_ = this.deleteCase_.bind(this);
    }

    deleteCase_(id){
        Case_Service.deleteCase_(id).then( res => {
            this.setState({case_s: this.state.case_s.filter(case_ => case_.case_Id !== id)});
        });
    }
    viewCase_(id){
        this.props.history.push(`/view-case_/${id}`);
    }
    editCase_(id){
        this.props.history.push(`/add-case_/${id}`);
    }

    componentDidMount(){
        Case_Service.getCase_s().then((res) => {
            this.setState({ case_s: res.data});
        });
    }

    addCase_(){
        this.props.history.push('/add-case_/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Case_ List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCase_}> Add Case_</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> CaseNumber </th>
                                    <th> Subject </th>
                                    <th> Description </th>
                                    <th> SlaDue </th>
                                    <th> Status </th>
                                    <th> Priority </th>
                                    <th> Origin </th>
                                    <th> Severity </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.case_s.map(
                                        case_ => 
                                        <tr key = {case_.case_Id}>
                                             <td> { case_.caseNumber } </td>
                                             <td> { case_.subject } </td>
                                             <td> { case_.description } </td>
                                             <td> { case_.slaDue } </td>
                                             <td> { case_.status } </td>
                                             <td> { case_.priority } </td>
                                             <td> { case_.origin } </td>
                                             <td> { case_.severity } </td>
                                             <td>
                                                 <button onClick={ () => this.editCase_(case_.case_Id)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCase_(case_.case_Id)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCase_(case_.case_Id)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListCase_Component
