
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreatePlacementComponent } from './create.component';
import { PlacementService } from '../../../services/Placement.service';
import { Router } from '@angular/router';

describe('CreatePlacementComponent', () => {
  let component: CreatePlacementComponent;
  let fixture: ComponentFixture<CreatePlacementComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreatePlacementComponent
      ],
      providers: [
        PlacementService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreatePlacementComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});