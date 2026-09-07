
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateAPUComponent } from './create.component';
import { APUService } from '../../../services/APU.service';
import { Router } from '@angular/router';

describe('CreateAPUComponent', () => {
  let component: CreateAPUComponent;
  let fixture: ComponentFixture<CreateAPUComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateAPUComponent
      ],
      providers: [
        APUService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateAPUComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});