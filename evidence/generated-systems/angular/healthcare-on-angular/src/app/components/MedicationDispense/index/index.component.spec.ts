
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexMedicationDispenseComponent } from './index.component';
import { MedicationDispenseService } from '../../../services/MedicationDispense.service';

describe('IndexMedicationDispenseComponent', () => {
  let component: IndexMedicationDispenseComponent;
  let fixture: ComponentFixture<IndexMedicationDispenseComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexMedicationDispenseComponent
      ],
      providers: [
        MedicationDispenseService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexMedicationDispenseComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});