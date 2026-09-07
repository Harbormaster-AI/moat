
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { DataSetService } from '../../../services/DataSet.service';
import { DataSet } from '../../../models/DataSet';

@Component({
    selector: 'app-index-dataSet',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexDataSetComponent implements OnInit {

    dataSets: DataSet[] = [];

    constructor(
        private router: Router,
        private service: DataSetService
) {}

    ngOnInit(): void {
        this.getDataSets();
}

    getDataSets(): void {
        this.service.getDataSets().subscribe((res) => {
        this.dataSets = res;
    });
}

    deleteDataSet(id: any): void {
        this.service.deleteDataSet(id)
            .subscribe(() => {
                this.getDataSets();
            });
    }
}