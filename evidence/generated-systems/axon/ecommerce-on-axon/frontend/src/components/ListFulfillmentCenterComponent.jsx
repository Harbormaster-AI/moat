import React, { Component } from 'react'
import FulfillmentCenterService from '../services/FulfillmentCenterService'

class ListFulfillmentCenterComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                fulfillmentCenters: []
        }
        this.addFulfillmentCenter = this.addFulfillmentCenter.bind(this);
        this.editFulfillmentCenter = this.editFulfillmentCenter.bind(this);
        this.deleteFulfillmentCenter = this.deleteFulfillmentCenter.bind(this);
    }

    deleteFulfillmentCenter(id){
        FulfillmentCenterService.deleteFulfillmentCenter(id).then( res => {
            this.setState({fulfillmentCenters: this.state.fulfillmentCenters.filter(fulfillmentCenter => fulfillmentCenter.fulfillmentCenterId !== id)});
        });
    }
    viewFulfillmentCenter(id){
        this.props.history.push(`/view-fulfillmentCenter/${id}`);
    }
    editFulfillmentCenter(id){
        this.props.history.push(`/add-fulfillmentCenter/${id}`);
    }

    componentDidMount(){
        FulfillmentCenterService.getFulfillmentCenters().then((res) => {
            this.setState({ fulfillmentCenters: res.data});
        });
    }

    addFulfillmentCenter(){
        this.props.history.push('/add-fulfillmentCenter/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">FulfillmentCenter List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addFulfillmentCenter}> Add FulfillmentCenter</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> CenterCode </th>
                                    <th> Address </th>
                                    <th> Timezone </th>
                                    <th> AsActive </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.fulfillmentCenters.map(
                                        fulfillmentCenter => 
                                        <tr key = {fulfillmentCenter.fulfillmentCenterId}>
                                             <td> { fulfillmentCenter.name } </td>
                                             <td> { fulfillmentCenter.centerCode } </td>
                                             <td> { fulfillmentCenter.address } </td>
                                             <td> { fulfillmentCenter.timezone } </td>
                                             <td> { fulfillmentCenter.asActive } </td>
                                             <td>
                                                 <button onClick={ () => this.editFulfillmentCenter(fulfillmentCenter.fulfillmentCenterId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteFulfillmentCenter(fulfillmentCenter.fulfillmentCenterId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewFulfillmentCenter(fulfillmentCenter.fulfillmentCenterId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListFulfillmentCenterComponent
