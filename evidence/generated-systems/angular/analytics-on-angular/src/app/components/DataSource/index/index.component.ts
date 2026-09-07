
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { DataSourceService } from '../../../services/DataSource.service';
import { DataSource } from '../../../models/DataSource';

@Component({
    selector: 'app-index-dataSource',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexDataSourceComponent implements OnInit {

    dataSources: DataSource[] = [];

    constructor(
        private router: Router,
        private service: DataSourceService
) {}

    ngOnInit(): void {
        this.getDataSources();
}

    getDataSources(): void {
        this.service.getDataSources().subscribe((res) => {
        this.dataSources = res;
    });
}

    deleteDataSource(id: any): void {
        this.service.deleteDataSource(id)
            .subscribe(() => {
                this.getDataSources();
            });
    }
}