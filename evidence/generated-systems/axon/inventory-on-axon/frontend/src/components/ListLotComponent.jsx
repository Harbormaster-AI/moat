import React, { Component } from 'react'
import LotService from '../services/LotService'

class ListLotComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                lots: []
        }
        this.addLot = this.addLot.bind(this);
        this.editLot = this.editLot.bind(this);
        this.deleteLot = this.deleteLot.bind(this);
    }

    deleteLot(id){
        LotService.deleteLot(id).then( res => {
            this.setState({lots: this.state.lots.filter(lot => lot.lotId !== id)});
        });
    }
    viewLot(id){
        this.props.history.push(`/view-lot/${id}`);
    }
    editLot(id){
        this.props.history.push(`/add-lot/${id}`);
    }

    componentDidMount(){
        LotService.getLots().then((res) => {
            this.setState({ lots: res.data});
        });
    }

    addLot(){
        this.props.history.push('/add-lot/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Lot List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addLot}> Add Lot</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> BatchNumber </th>
                                    <th> ManufactureDate </th>
                                    <th> ExpirationDate </th>
                                    <th> LotStatus </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.lots.map(
                                        lot => 
                                        <tr key = {lot.lotId}>
                                             <td> { lot.batchNumber } </td>
                                             <td> { lot.manufactureDate } </td>
                                             <td> { lot.expirationDate } </td>
                                             <td> { lot.lotStatus } </td>
                                             <td>
                                                 <button onClick={ () => this.editLot(lot.lotId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteLot(lot.lotId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewLot(lot.lotId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListLotComponent
