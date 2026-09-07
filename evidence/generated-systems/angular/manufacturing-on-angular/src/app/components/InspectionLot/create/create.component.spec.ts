
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateInspectionLotComponent } from './create.component';
import { InspectionLotService } from '../../../services/InspectionLot.service';
import { Router } from '@angular/router';

describe('CreateInspectionLotComponent', () => {
  let component: CreateInspectionLotComponent;
  let fixture: ComponentFixture<CreateInspectionLotComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateInspectionLotComponent
      ],
      providers: [
        InspectionLotService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateInspectionLotComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});