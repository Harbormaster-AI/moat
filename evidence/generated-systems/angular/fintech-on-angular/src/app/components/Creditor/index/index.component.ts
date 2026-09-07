
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CreditorService } from '../../../services/Creditor.service';
import { Creditor } from '../../../models/Creditor';

@Component({
    selector: 'app-index-creditor',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexCreditorComponent implements OnInit {

    creditors: Creditor[] = [];

    constructor(
        private router: Router,
        private service: CreditorService
) {}

    ngOnInit(): void {
        this.getCreditors();
}

    getCreditors(): void {
        this.service.getCreditors().subscribe((res) => {
        this.creditors = res;
    });
}

    deleteCreditor(id: any): void {
        this.service.deleteCreditor(id)
            .subscribe(() => {
                this.getCreditors();
            });
    }
}