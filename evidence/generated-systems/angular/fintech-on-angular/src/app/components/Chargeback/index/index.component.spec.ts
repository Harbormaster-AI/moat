
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexChargebackComponent } from './index.component';
import { ChargebackService } from '../../../services/Chargeback.service';

describe('IndexChargebackComponent', () => {
  let component: IndexChargebackComponent;
  let fixture: ComponentFixture<IndexChargebackComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexChargebackComponent
      ],
      providers: [
        ChargebackService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexChargebackComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});