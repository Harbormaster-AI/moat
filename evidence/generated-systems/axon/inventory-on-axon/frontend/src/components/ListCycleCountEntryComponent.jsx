import React, { Component } from 'react'
import CycleCountEntryService from '../services/CycleCountEntryService'

class ListCycleCountEntryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                cycleCountEntrys: []
        }
        this.addCycleCountEntry = this.addCycleCountEntry.bind(this);
        this.editCycleCountEntry = this.editCycleCountEntry.bind(this);
        this.deleteCycleCountEntry = this.deleteCycleCountEntry.bind(this);
    }

    deleteCycleCountEntry(id){
        CycleCountEntryService.deleteCycleCountEntry(id).then( res => {
            this.setState({cycleCountEntrys: this.state.cycleCountEntrys.filter(cycleCountEntry => cycleCountEntry.cycleCountEntryId !== id)});
        });
    }
    viewCycleCountEntry(id){
        this.props.history.push(`/view-cycleCountEntry/${id}`);
    }
    editCycleCountEntry(id){
        this.props.history.push(`/add-cycleCountEntry/${id}`);
    }

    componentDidMount(){
        CycleCountEntryService.getCycleCountEntrys().then((res) => {
            this.setState({ cycleCountEntrys: res.data});
        });
    }

    addCycleCountEntry(){
        this.props.history.push('/add-cycleCountEntry/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">CycleCountEntry List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCycleCountEntry}> Add CycleCountEntry</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> LineNumber </th>
                                    <th> SystemQuantity </th>
                                    <th> CountedQuantity </th>
                                    <th> VarianceQuantity </th>
                                    <th> RecountRequired </th>
                                    <th> StockStatus </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.cycleCountEntrys.map(
                                        cycleCountEntry => 
                                        <tr key = {cycleCountEntry.cycleCountEntryId}>
                                             <td> { cycleCountEntry.lineNumber } </td>
                                             <td> { cycleCountEntry.systemQuantity } </td>
                                             <td> { cycleCountEntry.countedQuantity } </td>
                                             <td> { cycleCountEntry.varianceQuantity } </td>
                                             <td> { cycleCountEntry.recountRequired } </td>
                                             <td> { cycleCountEntry.stockStatus } </td>
                                             <td>
                                                 <button onClick={ () => this.editCycleCountEntry(cycleCountEntry.cycleCountEntryId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCycleCountEntry(cycleCountEntry.cycleCountEntryId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCycleCountEntry(cycleCountEntry.cycleCountEntryId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListCycleCountEntryComponent
