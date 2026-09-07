
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { DataProviderService } from '../../../services/DataProvider.service';
import { DataProvider } from '../../../models/DataProvider';

@Component({
    selector: 'app-index-dataProvider',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexDataProviderComponent implements OnInit {

    dataProviders: DataProvider[] = [];

    constructor(
        private router: Router,
        private service: DataProviderService
) {}

    ngOnInit(): void {
        this.getDataProviders();
}

    getDataProviders(): void {
        this.service.getDataProviders().subscribe((res) => {
        this.dataProviders = res;
    });
}

    deleteDataProvider(id: any): void {
        this.service.deleteDataProvider(id)
            .subscribe(() => {
                this.getDataProviders();
            });
    }
}