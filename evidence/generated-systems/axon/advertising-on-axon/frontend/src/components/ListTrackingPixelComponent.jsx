import React, { Component } from 'react'
import TrackingPixelService from '../services/TrackingPixelService'

class ListTrackingPixelComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                trackingPixels: []
        }
        this.addTrackingPixel = this.addTrackingPixel.bind(this);
        this.editTrackingPixel = this.editTrackingPixel.bind(this);
        this.deleteTrackingPixel = this.deleteTrackingPixel.bind(this);
    }

    deleteTrackingPixel(id){
        TrackingPixelService.deleteTrackingPixel(id).then( res => {
            this.setState({trackingPixels: this.state.trackingPixels.filter(trackingPixel => trackingPixel.trackingPixelId !== id)});
        });
    }
    viewTrackingPixel(id){
        this.props.history.push(`/view-trackingPixel/${id}`);
    }
    editTrackingPixel(id){
        this.props.history.push(`/add-trackingPixel/${id}`);
    }

    componentDidMount(){
        TrackingPixelService.getTrackingPixels().then((res) => {
            this.setState({ trackingPixels: res.data});
        });
    }

    addTrackingPixel(){
        this.props.history.push('/add-trackingPixel/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">TrackingPixel List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addTrackingPixel}> Add TrackingPixel</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Url </th>
                                    <th> EventType </th>
                                    <th> PixelType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.trackingPixels.map(
                                        trackingPixel => 
                                        <tr key = {trackingPixel.trackingPixelId}>
                                             <td> { trackingPixel.name } </td>
                                             <td> { trackingPixel.url } </td>
                                             <td> { trackingPixel.eventType } </td>
                                             <td> { trackingPixel.pixelType } </td>
                                             <td>
                                                 <button onClick={ () => this.editTrackingPixel(trackingPixel.trackingPixelId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteTrackingPixel(trackingPixel.trackingPixelId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewTrackingPixel(trackingPixel.trackingPixelId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListTrackingPixelComponent
