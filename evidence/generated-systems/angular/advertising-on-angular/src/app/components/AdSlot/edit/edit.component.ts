import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AdSlotService } from '../../../services/AdSlot.service';
import { SubBaseComponent } from '../../AdSlot/sub.base.component';


@Component({
    selector: 'app-edit-adSlot',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAdSlotComponent extends SubBaseComponent implements OnInit {

    title = 'Edit AdSlot';

    adSlotForm: FormGroup;
    adSlot: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AdSlotService,
        private fb: FormBuilder
) {
        super(http);
        this.adSlotForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  slotCode: ['', Validators.required],
      width: ['', Validators.required],
      height: ['', Validators.required],
      floorPrice: ['', Validators.required],
      InventorySource: ['', ],
      Placements: ['', ],
      Rates: ['', ],
      Format: ['', ]
        });
    }

    
    updateAdSlot(slotCode, width, height, floorPrice, InventorySource, Placements, Rates, Format): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAdSlot(slotCode, width, height, floorPrice, InventorySource, Placements, Rates, Format, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAdSlot']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAdSlot(params['id']).subscribe(res => {
                this.adSlot = res;
            });
        });
    }
}