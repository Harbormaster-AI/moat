import React, { Component } from 'react'
import DataSourceService from '../services/DataSourceService'

class ViewDataSourceComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            dataSource: {}
        }
    }

    componentDidMount(){
        DataSourceService.getDataSourceById(this.state.id).then( res => {
            this.setState({dataSource: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View DataSource Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dataSource.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> connection:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dataSource.connection }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Streaming:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dataSource.streaming }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> SourceType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dataSource.sourceType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Format:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dataSource.format }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewDataSourceComponent
