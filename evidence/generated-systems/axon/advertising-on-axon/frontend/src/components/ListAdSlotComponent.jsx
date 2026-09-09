import React, { Component } from 'react'
import AdSlotService from '../services/AdSlotService'

class ListAdSlotComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                adSlots: []
        }
        this.addAdSlot = this.addAdSlot.bind(this);
        this.editAdSlot = this.editAdSlot.bind(this);
        this.deleteAdSlot = this.deleteAdSlot.bind(this);
    }

    deleteAdSlot(id){
        AdSlotService.deleteAdSlot(id).then( res => {
            this.setState({adSlots: this.state.adSlots.filter(adSlot => adSlot.adSlotId !== id)});
        });
    }
    viewAdSlot(id){
        this.props.history.push(`/view-adSlot/${id}`);
    }
    editAdSlot(id){
        this.props.history.push(`/add-adSlot/${id}`);
    }

    componentDidMount(){
        AdSlotService.getAdSlots().then((res) => {
            this.setState({ adSlots: res.data});
        });
    }

    addAdSlot(){
        this.props.history.push('/add-adSlot/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">AdSlot List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAdSlot}> Add AdSlot</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> SlotCode </th>
                                    <th> Width </th>
                                    <th> Height </th>
                                    <th> FloorPrice </th>
                                    <th> Format </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.adSlots.map(
                                        adSlot => 
                                        <tr key = {adSlot.adSlotId}>
                                             <td> { adSlot.slotCode } </td>
                                             <td> { adSlot.width } </td>
                                             <td> { adSlot.height } </td>
                                             <td> { adSlot.floorPrice } </td>
                                             <td> { adSlot.format } </td>
                                             <td>
                                                 <button onClick={ () => this.editAdSlot(adSlot.adSlotId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAdSlot(adSlot.adSlotId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAdSlot(adSlot.adSlotId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAdSlotComponent
