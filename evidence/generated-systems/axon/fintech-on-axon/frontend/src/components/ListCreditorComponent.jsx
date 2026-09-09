import React, { Component } from 'react'
import CreditorService from '../services/CreditorService'

class ListCreditorComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                creditors: []
        }
        this.addCreditor = this.addCreditor.bind(this);
        this.editCreditor = this.editCreditor.bind(this);
        this.deleteCreditor = this.deleteCreditor.bind(this);
    }

    deleteCreditor(id){
        CreditorService.deleteCreditor(id).then( res => {
            this.setState({creditors: this.state.creditors.filter(creditor => creditor.creditorId !== id)});
        });
    }
    viewCreditor(id){
        this.props.history.push(`/view-creditor/${id}`);
    }
    editCreditor(id){
        this.props.history.push(`/add-creditor/${id}`);
    }

    componentDidMount(){
        CreditorService.getCreditors().then((res) => {
            this.setState({ creditors: res.data});
        });
    }

    addCreditor(){
        this.props.history.push('/add-creditor/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Creditor List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCreditor}> Add Creditor</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Bic </th>
                                    <th> Address </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.creditors.map(
                                        creditor => 
                                        <tr key = {creditor.creditorId}>
                                             <td> { creditor.name } </td>
                                             <td> { creditor.bic } </td>
                                             <td> { creditor.address } </td>
                                             <td>
                                                 <button onClick={ () => this.editCreditor(creditor.creditorId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCreditor(creditor.creditorId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCreditor(creditor.creditorId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListCreditorComponent
