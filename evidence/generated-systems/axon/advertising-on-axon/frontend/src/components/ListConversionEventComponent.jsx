import React, { Component } from 'react'
import ConversionEventService from '../services/ConversionEventService'

class ListConversionEventComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                conversionEvents: []
        }
        this.addConversionEvent = this.addConversionEvent.bind(this);
        this.editConversionEvent = this.editConversionEvent.bind(this);
        this.deleteConversionEvent = this.deleteConversionEvent.bind(this);
    }

    deleteConversionEvent(id){
        ConversionEventService.deleteConversionEvent(id).then( res => {
            this.setState({conversionEvents: this.state.conversionEvents.filter(conversionEvent => conversionEvent.conversionEventId !== id)});
        });
    }
    viewConversionEvent(id){
        this.props.history.push(`/view-conversionEvent/${id}`);
    }
    editConversionEvent(id){
        this.props.history.push(`/add-conversionEvent/${id}`);
    }

    componentDidMount(){
        ConversionEventService.getConversionEvents().then((res) => {
            this.setState({ conversionEvents: res.data});
        });
    }

    addConversionEvent(){
        this.props.history.push('/add-conversionEvent/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ConversionEvent List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addConversionEvent}> Add ConversionEvent</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Timestamp </th>
                                    <th> Value </th>
                                    <th> EventType </th>
                                    <th> AttributionModel </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.conversionEvents.map(
                                        conversionEvent => 
                                        <tr key = {conversionEvent.conversionEventId}>
                                             <td> { conversionEvent.timestamp } </td>
                                             <td> { conversionEvent.value } </td>
                                             <td> { conversionEvent.eventType } </td>
                                             <td> { conversionEvent.attributionModel } </td>
                                             <td>
                                                 <button onClick={ () => this.editConversionEvent(conversionEvent.conversionEventId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteConversionEvent(conversionEvent.conversionEventId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewConversionEvent(conversionEvent.conversionEventId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListConversionEventComponent
