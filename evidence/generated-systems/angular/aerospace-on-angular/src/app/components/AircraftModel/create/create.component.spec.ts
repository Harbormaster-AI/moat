
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateAircraftModelComponent } from './create.component';
import { AircraftModelService } from '../../../services/AircraftModel.service';
import { Router } from '@angular/router';

describe('CreateAircraftModelComponent', () => {
  let component: CreateAircraftModelComponent;
  let fixture: ComponentFixture<CreateAircraftModelComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateAircraftModelComponent
      ],
      providers: [
        AircraftModelService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateAircraftModelComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});