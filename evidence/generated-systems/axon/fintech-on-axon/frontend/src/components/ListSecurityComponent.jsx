import React, { Component } from 'react'
import SecurityService from '../services/SecurityService'

class ListSecurityComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                securitys: []
        }
        this.addSecurity = this.addSecurity.bind(this);
        this.editSecurity = this.editSecurity.bind(this);
        this.deleteSecurity = this.deleteSecurity.bind(this);
    }

    deleteSecurity(id){
        SecurityService.deleteSecurity(id).then( res => {
            this.setState({securitys: this.state.securitys.filter(security => security.securityId !== id)});
        });
    }
    viewSecurity(id){
        this.props.history.push(`/view-security/${id}`);
    }
    editSecurity(id){
        this.props.history.push(`/add-security/${id}`);
    }

    componentDidMount(){
        SecurityService.getSecuritys().then((res) => {
            this.setState({ securitys: res.data});
        });
    }

    addSecurity(){
        this.props.history.push('/add-security/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Security List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addSecurity}> Add Security</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Symbol </th>
                                    <th> Isin </th>
                                    <th> Cusip </th>
                                    <th> Currency </th>
                                    <th> SecurityType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.securitys.map(
                                        security => 
                                        <tr key = {security.securityId}>
                                             <td> { security.symbol } </td>
                                             <td> { security.isin } </td>
                                             <td> { security.cusip } </td>
                                             <td> { security.currency } </td>
                                             <td> { security.securityType } </td>
                                             <td>
                                                 <button onClick={ () => this.editSecurity(security.securityId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteSecurity(security.securityId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewSecurity(security.securityId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListSecurityComponent
