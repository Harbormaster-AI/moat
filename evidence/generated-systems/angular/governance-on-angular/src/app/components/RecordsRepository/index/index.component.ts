
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { RecordsRepositoryService } from '../../../services/RecordsRepository.service';
import { RecordsRepository } from '../../../models/RecordsRepository';

@Component({
    selector: 'app-index-recordsRepository',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexRecordsRepositoryComponent implements OnInit {

    recordsRepositorys: RecordsRepository[] = [];

    constructor(
        private router: Router,
        private service: RecordsRepositoryService
) {}

    ngOnInit(): void {
        this.getRecordsRepositorys();
}

    getRecordsRepositorys(): void {
        this.service.getRecordsRepositorys().subscribe((res) => {
        this.recordsRepositorys = res;
    });
}

    deleteRecordsRepository(id: any): void {
        this.service.deleteRecordsRepository(id)
            .subscribe(() => {
                this.getRecordsRepositorys();
            });
    }
}