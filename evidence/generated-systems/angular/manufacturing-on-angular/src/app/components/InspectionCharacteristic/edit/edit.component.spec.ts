
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditInspectionCharacteristicComponent } from './edit.component';
import { InspectionCharacteristicService } from '../../../services/InspectionCharacteristic.service';

describe('EditInspectionCharacteristicComponent', () => {
  let component: EditInspectionCharacteristicComponent;
  let fixture: ComponentFixture<EditInspectionCharacteristicComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditInspectionCharacteristicComponent
      ],
      providers: [
        InspectionCharacteristicService,
        {
          provide: ActivatedRoute,
          useValue: {
            params: of({ id: '1' })
          }
        },
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(EditInspectionCharacteristicComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});