import React, { Component } from 'react'
import RoutingService from '../services/RoutingService'

class ListRoutingComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                routings: []
        }
        this.addRouting = this.addRouting.bind(this);
        this.editRouting = this.editRouting.bind(this);
        this.deleteRouting = this.deleteRouting.bind(this);
    }

    deleteRouting(id){
        RoutingService.deleteRouting(id).then( res => {
            this.setState({routings: this.state.routings.filter(routing => routing.routingId !== id)});
        });
    }
    viewRouting(id){
        this.props.history.push(`/view-routing/${id}`);
    }
    editRouting(id){
        this.props.history.push(`/add-routing/${id}`);
    }

    componentDidMount(){
        RoutingService.getRoutings().then((res) => {
            this.setState({ routings: res.data});
        });
    }

    addRouting(){
        this.props.history.push('/add-routing/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Routing List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addRouting}> Add Routing</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> RoutingNumber </th>
                                    <th> Revision </th>
                                    <th> EffectivityStart </th>
                                    <th> EffectivityEnd </th>
                                    <th> RoutingType </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.routings.map(
                                        routing => 
                                        <tr key = {routing.routingId}>
                                             <td> { routing.routingNumber } </td>
                                             <td> { routing.revision } </td>
                                             <td> { routing.effectivityStart } </td>
                                             <td> { routing.effectivityEnd } </td>
                                             <td> { routing.routingType } </td>
                                             <td> { routing.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editRouting(routing.routingId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteRouting(routing.routingId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewRouting(routing.routingId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListRoutingComponent
