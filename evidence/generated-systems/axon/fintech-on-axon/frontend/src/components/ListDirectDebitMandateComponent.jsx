import React, { Component } from 'react'
import DirectDebitMandateService from '../services/DirectDebitMandateService'

class ListDirectDebitMandateComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                directDebitMandates: []
        }
        this.addDirectDebitMandate = this.addDirectDebitMandate.bind(this);
        this.editDirectDebitMandate = this.editDirectDebitMandate.bind(this);
        this.deleteDirectDebitMandate = this.deleteDirectDebitMandate.bind(this);
    }

    deleteDirectDebitMandate(id){
        DirectDebitMandateService.deleteDirectDebitMandate(id).then( res => {
            this.setState({directDebitMandates: this.state.directDebitMandates.filter(directDebitMandate => directDebitMandate.directDebitMandateId !== id)});
        });
    }
    viewDirectDebitMandate(id){
        this.props.history.push(`/view-directDebitMandate/${id}`);
    }
    editDirectDebitMandate(id){
        this.props.history.push(`/add-directDebitMandate/${id}`);
    }

    componentDidMount(){
        DirectDebitMandateService.getDirectDebitMandates().then((res) => {
            this.setState({ directDebitMandates: res.data});
        });
    }

    addDirectDebitMandate(){
        this.props.history.push('/add-directDebitMandate/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">DirectDebitMandate List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addDirectDebitMandate}> Add DirectDebitMandate</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> MandateId </th>
                                    <th> SignedAt </th>
                                    <th> Scheme </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.directDebitMandates.map(
                                        directDebitMandate => 
                                        <tr key = {directDebitMandate.directDebitMandateId}>
                                             <td> { directDebitMandate.mandateId } </td>
                                             <td> { directDebitMandate.signedAt } </td>
                                             <td> { directDebitMandate.scheme } </td>
                                             <td> { directDebitMandate.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editDirectDebitMandate(directDebitMandate.directDebitMandateId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteDirectDebitMandate(directDebitMandate.directDebitMandateId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewDirectDebitMandate(directDebitMandate.directDebitMandateId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListDirectDebitMandateComponent
