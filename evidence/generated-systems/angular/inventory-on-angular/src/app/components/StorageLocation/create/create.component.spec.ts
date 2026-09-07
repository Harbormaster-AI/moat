
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateStorageLocationComponent } from './create.component';
import { StorageLocationService } from '../../../services/StorageLocation.service';
import { Router } from '@angular/router';

describe('CreateStorageLocationComponent', () => {
  let component: CreateStorageLocationComponent;
  let fixture: ComponentFixture<CreateStorageLocationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateStorageLocationComponent
      ],
      providers: [
        StorageLocationService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateStorageLocationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});