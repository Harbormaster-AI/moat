
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { DataBreachService } from '../../../services/DataBreach.service';
import { DataBreach } from '../../../models/DataBreach';

@Component({
    selector: 'app-index-dataBreach',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexDataBreachComponent implements OnInit {

    dataBreachs: DataBreach[] = [];

    constructor(
        private router: Router,
        private service: DataBreachService
) {}

    ngOnInit(): void {
        this.getDataBreachs();
}

    getDataBreachs(): void {
        this.service.getDataBreachs().subscribe((res) => {
        this.dataBreachs = res;
    });
}

    deleteDataBreach(id: any): void {
        this.service.deleteDataBreach(id)
            .subscribe(() => {
                this.getDataBreachs();
            });
    }
}