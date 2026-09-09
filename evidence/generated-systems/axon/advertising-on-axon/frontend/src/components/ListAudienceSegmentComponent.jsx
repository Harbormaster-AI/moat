import React, { Component } from 'react'
import AudienceSegmentService from '../services/AudienceSegmentService'

class ListAudienceSegmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                audienceSegments: []
        }
        this.addAudienceSegment = this.addAudienceSegment.bind(this);
        this.editAudienceSegment = this.editAudienceSegment.bind(this);
        this.deleteAudienceSegment = this.deleteAudienceSegment.bind(this);
    }

    deleteAudienceSegment(id){
        AudienceSegmentService.deleteAudienceSegment(id).then( res => {
            this.setState({audienceSegments: this.state.audienceSegments.filter(audienceSegment => audienceSegment.audienceSegmentId !== id)});
        });
    }
    viewAudienceSegment(id){
        this.props.history.push(`/view-audienceSegment/${id}`);
    }
    editAudienceSegment(id){
        this.props.history.push(`/add-audienceSegment/${id}`);
    }

    componentDidMount(){
        AudienceSegmentService.getAudienceSegments().then((res) => {
            this.setState({ audienceSegments: res.data});
        });
    }

    addAudienceSegment(){
        this.props.history.push('/add-audienceSegment/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">AudienceSegment List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAudienceSegment}> Add AudienceSegment</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> EstimatedReach </th>
                                    <th> Description </th>
                                    <th> ProviderType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.audienceSegments.map(
                                        audienceSegment => 
                                        <tr key = {audienceSegment.audienceSegmentId}>
                                             <td> { audienceSegment.name } </td>
                                             <td> { audienceSegment.estimatedReach } </td>
                                             <td> { audienceSegment.description } </td>
                                             <td> { audienceSegment.providerType } </td>
                                             <td>
                                                 <button onClick={ () => this.editAudienceSegment(audienceSegment.audienceSegmentId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAudienceSegment(audienceSegment.audienceSegmentId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAudienceSegment(audienceSegment.audienceSegmentId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAudienceSegmentComponent
