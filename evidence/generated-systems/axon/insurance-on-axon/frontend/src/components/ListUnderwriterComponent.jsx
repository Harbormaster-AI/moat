import React, { Component } from 'react'
import UnderwriterService from '../services/UnderwriterService'

class ListUnderwriterComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                underwriters: []
        }
        this.addUnderwriter = this.addUnderwriter.bind(this);
        this.editUnderwriter = this.editUnderwriter.bind(this);
        this.deleteUnderwriter = this.deleteUnderwriter.bind(this);
    }

    deleteUnderwriter(id){
        UnderwriterService.deleteUnderwriter(id).then( res => {
            this.setState({underwriters: this.state.underwriters.filter(underwriter => underwriter.underwriterId !== id)});
        });
    }
    viewUnderwriter(id){
        this.props.history.push(`/view-underwriter/${id}`);
    }
    editUnderwriter(id){
        this.props.history.push(`/add-underwriter/${id}`);
    }

    componentDidMount(){
        UnderwriterService.getUnderwriters().then((res) => {
            this.setState({ underwriters: res.data});
        });
    }

    addUnderwriter(){
        this.props.history.push('/add-underwriter/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Underwriter List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addUnderwriter}> Add Underwriter</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> FirstName </th>
                                    <th> LastName </th>
                                    <th> EmployeeId </th>
                                    <th> AuthorityLimit </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.underwriters.map(
                                        underwriter => 
                                        <tr key = {underwriter.underwriterId}>
                                             <td> { underwriter.firstName } </td>
                                             <td> { underwriter.lastName } </td>
                                             <td> { underwriter.employeeId } </td>
                                             <td> { underwriter.authorityLimit } </td>
                                             <td>
                                                 <button onClick={ () => this.editUnderwriter(underwriter.underwriterId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteUnderwriter(underwriter.underwriterId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewUnderwriter(underwriter.underwriterId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListUnderwriterComponent
