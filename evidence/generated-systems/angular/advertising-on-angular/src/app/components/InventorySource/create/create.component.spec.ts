
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateInventorySourceComponent } from './create.component';
import { InventorySourceService } from '../../../services/InventorySource.service';
import { Router } from '@angular/router';

describe('CreateInventorySourceComponent', () => {
  let component: CreateInventorySourceComponent;
  let fixture: ComponentFixture<CreateInventorySourceComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateInventorySourceComponent
      ],
      providers: [
        InventorySourceService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateInventorySourceComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});