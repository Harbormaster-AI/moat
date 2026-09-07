
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { DataSubjectRequestService } from '../../../services/DataSubjectRequest.service';
import { DataSubjectRequest } from '../../../models/DataSubjectRequest';

@Component({
    selector: 'app-index-dataSubjectRequest',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexDataSubjectRequestComponent implements OnInit {

    dataSubjectRequests: DataSubjectRequest[] = [];

    constructor(
        private router: Router,
        private service: DataSubjectRequestService
) {}

    ngOnInit(): void {
        this.getDataSubjectRequests();
}

    getDataSubjectRequests(): void {
        this.service.getDataSubjectRequests().subscribe((res) => {
        this.dataSubjectRequests = res;
    });
}

    deleteDataSubjectRequest(id: any): void {
        this.service.deleteDataSubjectRequest(id)
            .subscribe(() => {
                this.getDataSubjectRequests();
            });
    }
}