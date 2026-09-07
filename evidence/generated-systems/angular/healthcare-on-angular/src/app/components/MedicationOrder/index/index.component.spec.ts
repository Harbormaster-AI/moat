
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexMedicationOrderComponent } from './index.component';
import { MedicationOrderService } from '../../../services/MedicationOrder.service';

describe('IndexMedicationOrderComponent', () => {
  let component: IndexMedicationOrderComponent;
  let fixture: ComponentFixture<IndexMedicationOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexMedicationOrderComponent
      ],
      providers: [
        MedicationOrderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexMedicationOrderComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});