
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateAirworthinessDirectiveComponent } from './create.component';
import { AirworthinessDirectiveService } from '../../../services/AirworthinessDirective.service';
import { Router } from '@angular/router';

describe('CreateAirworthinessDirectiveComponent', () => {
  let component: CreateAirworthinessDirectiveComponent;
  let fixture: ComponentFixture<CreateAirworthinessDirectiveComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateAirworthinessDirectiveComponent
      ],
      providers: [
        AirworthinessDirectiveService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateAirworthinessDirectiveComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});