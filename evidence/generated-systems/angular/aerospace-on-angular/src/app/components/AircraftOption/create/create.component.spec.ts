
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateAircraftOptionComponent } from './create.component';
import { AircraftOptionService } from '../../../services/AircraftOption.service';
import { Router } from '@angular/router';

describe('CreateAircraftOptionComponent', () => {
  let component: CreateAircraftOptionComponent;
  let fixture: ComponentFixture<CreateAircraftOptionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateAircraftOptionComponent
      ],
      providers: [
        AircraftOptionService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateAircraftOptionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});