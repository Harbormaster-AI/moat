import React, { Component } from 'react'
import ServiceBulletinService from '../services/ServiceBulletinService'

class ListServiceBulletinComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                serviceBulletins: []
        }
        this.addServiceBulletin = this.addServiceBulletin.bind(this);
        this.editServiceBulletin = this.editServiceBulletin.bind(this);
        this.deleteServiceBulletin = this.deleteServiceBulletin.bind(this);
    }

    deleteServiceBulletin(id){
        ServiceBulletinService.deleteServiceBulletin(id).then( res => {
            this.setState({serviceBulletins: this.state.serviceBulletins.filter(serviceBulletin => serviceBulletin.serviceBulletinId !== id)});
        });
    }
    viewServiceBulletin(id){
        this.props.history.push(`/view-serviceBulletin/${id}`);
    }
    editServiceBulletin(id){
        this.props.history.push(`/add-serviceBulletin/${id}`);
    }

    componentDidMount(){
        ServiceBulletinService.getServiceBulletins().then((res) => {
            this.setState({ serviceBulletins: res.data});
        });
    }

    addServiceBulletin(){
        this.props.history.push('/add-serviceBulletin/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ServiceBulletin List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addServiceBulletin}> Add ServiceBulletin</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> BulletinNumber </th>
                                    <th> Category </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.serviceBulletins.map(
                                        serviceBulletin => 
                                        <tr key = {serviceBulletin.serviceBulletinId}>
                                             <td> { serviceBulletin.bulletinNumber } </td>
                                             <td> { serviceBulletin.category } </td>
                                             <td>
                                                 <button onClick={ () => this.editServiceBulletin(serviceBulletin.serviceBulletinId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteServiceBulletin(serviceBulletin.serviceBulletinId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewServiceBulletin(serviceBulletin.serviceBulletinId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListServiceBulletinComponent
