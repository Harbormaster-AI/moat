
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexMerchantComponent } from './index.component';
import { MerchantService } from '../../../services/Merchant.service';

describe('IndexMerchantComponent', () => {
  let component: IndexMerchantComponent;
  let fixture: ComponentFixture<IndexMerchantComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexMerchantComponent
      ],
      providers: [
        MerchantService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexMerchantComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});