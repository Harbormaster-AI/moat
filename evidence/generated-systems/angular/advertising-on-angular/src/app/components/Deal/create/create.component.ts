import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { DealService } from '../../../services/Deal.service';
import { Deal } from '../../../models/Deal';
import { SubBaseComponent } from '../../Deal/sub.base.component';

@Component({
    selector: 'app-create-deal',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateDealComponent extends SubBaseComponent implements OnInit {

    title = 'Add Deal';

    dealForm: FormGroup;
    deal: Deal;

    constructor( http: HttpClient,
        private dealService: DealService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.dealForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  floorPrice: ['', Validators.required],
      Publisher: ['', ],
      InventorySources: ['', ],
      Placements: ['', ],
      DealType: ['', ]
        });
    }

    
    addDeal(floorPrice, Publisher, InventorySources, Placements, DealType): void {
        this.dealService
        .addDeal(floorPrice, Publisher, InventorySources, Placements, DealType)
            .subscribe(() => {
                this.router.navigate(['/indexDeal']);
            });
    }

    ngOnInit(): void {
    }
}