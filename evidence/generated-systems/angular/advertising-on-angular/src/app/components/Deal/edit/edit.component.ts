import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { DealService } from '../../../services/Deal.service';
import { SubBaseComponent } from '../../Deal/sub.base.component';


@Component({
    selector: 'app-edit-deal',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditDealComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Deal';

    dealForm: FormGroup;
    deal: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: DealService,
        private fb: FormBuilder
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

    
    updateDeal(floorPrice, Publisher, InventorySources, Placements, DealType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateDeal(floorPrice, Publisher, InventorySources, Placements, DealType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexDeal']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getDeal(params['id']).subscribe(res => {
                this.deal = res;
            });
        });
    }
}