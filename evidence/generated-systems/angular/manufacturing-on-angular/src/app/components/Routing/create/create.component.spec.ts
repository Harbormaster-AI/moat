
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateRoutingComponent } from './create.component';
import { RoutingService } from '../../../services/Routing.service';
import { Router } from '@angular/router';

describe('CreateRoutingComponent', () => {
  let component: CreateRoutingComponent;
  let fixture: ComponentFixture<CreateRoutingComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateRoutingComponent
      ],
      providers: [
        RoutingService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateRoutingComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});