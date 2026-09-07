
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexRoutingComponent } from './index.component';
import { RoutingService } from '../../../services/Routing.service';

describe('IndexRoutingComponent', () => {
  let component: IndexRoutingComponent;
  let fixture: ComponentFixture<IndexRoutingComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexRoutingComponent
      ],
      providers: [
        RoutingService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexRoutingComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});