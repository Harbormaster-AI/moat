
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { DataPipelineService } from '../../../services/DataPipeline.service';
import { DataPipeline } from '../../../models/DataPipeline';

@Component({
    selector: 'app-index-dataPipeline',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexDataPipelineComponent implements OnInit {

    dataPipelines: DataPipeline[] = [];

    constructor(
        private router: Router,
        private service: DataPipelineService
) {}

    ngOnInit(): void {
        this.getDataPipelines();
}

    getDataPipelines(): void {
        this.service.getDataPipelines().subscribe((res) => {
        this.dataPipelines = res;
    });
}

    deleteDataPipeline(id: any): void {
        this.service.deleteDataPipeline(id)
            .subscribe(() => {
                this.getDataPipelines();
            });
    }
}