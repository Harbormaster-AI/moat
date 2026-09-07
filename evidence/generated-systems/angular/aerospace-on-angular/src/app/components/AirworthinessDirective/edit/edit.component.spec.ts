
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditAirworthinessDirectiveComponent } from './edit.component';
import { AirworthinessDirectiveService } from '../../../services/AirworthinessDirective.service';

describe('EditAirworthinessDirectiveComponent', () => {
  let component: EditAirworthinessDirectiveComponent;
  let fixture: ComponentFixture<EditAirworthinessDirectiveComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditAirworthinessDirectiveComponent
      ],
      providers: [
        AirworthinessDirectiveService,
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

    fixture = TestBed.createComponent(EditAirworthinessDirectiveComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});