
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AllergyService } from '../../../services/Allergy.service';
import { Allergy } from '../../../models/Allergy';

@Component({
    selector: 'app-index-allergy',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAllergyComponent implements OnInit {

    allergys: Allergy[] = [];

    constructor(
        private router: Router,
        private service: AllergyService
) {}

    ngOnInit(): void {
        this.getAllergys();
}

    getAllergys(): void {
        this.service.getAllergys().subscribe((res) => {
        this.allergys = res;
    });
}

    deleteAllergy(id: any): void {
        this.service.deleteAllergy(id)
            .subscribe(() => {
                this.getAllergys();
            });
    }
}